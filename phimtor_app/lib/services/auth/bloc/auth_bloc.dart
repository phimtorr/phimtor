import 'package:bloc/bloc.dart';
import 'package:phimtor_app/services/auth/auth_service.dart';
import 'package:phimtor_app/services/auth/bloc/auth_event.dart';
import 'package:phimtor_app/services/auth/bloc/auth_state.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  static AuthState getInitialState(AuthService authService) {
    final user = authService.currentUser;
    return user == null
        ? const AuthStateLoggedOut(exception: null, isLoading: false)
        : user.emailVerified == false
            ? const AuthStateNeedsVerification(isLoading: false)
            : AuthStateLoggedIn(user: user, isLoading: false);
  }

  AuthBloc(AuthService authService) : super(getInitialState(authService)) {
    // Handle AuthEventLogIn
    on<AuthEventLogIn>((event, emit) async {
      emit(const AuthStateLoggedOut(
        exception: null,
        isLoading: true,
        loadingText: 'Logging in...',
      ));

      final email = event.email;
      final password = event.password;

      try {
        final user = await authService.logIn(email: email, password: password);
        if (user.emailVerified == false) {
          emit(const AuthStateNeedsVerification(isLoading: false));
          return;
        }
        emit(AuthStateLoggedIn(user: user, isLoading: false));
      } catch (e) {
        emit(AuthStateLoggedOut(
          exception: e as Exception,
          isLoading: false,
        ));
      }
    });

    // Handle AuthEventLogOut
    on<AuthEventLogOut>((event, emit) async {
      emit(const AuthStateLoggedOut(
        exception: null,
        isLoading: true,
        loadingText: 'Logging out...',
      ));

      try {
        await authService.logOut();
        emit(const AuthStateLoggedOut(
          exception: null,
          isLoading: false,
        ));
      } on Exception catch (e) {
        emit(AuthStateLoggedOut(
          exception: e,
          isLoading: false,
        ));
      }
    });

    on<AuthEventShouldRegister>((event, emit) {
      emit(const AuthStateRegistering(exception: null, isLoading: false));
    });

    on<AuthEventReigister>((event, emit) async {
      final email = event.email;
      final password = event.password;

      try {
        await authService.createUser(email: email, password: password);
        await authService.sendEmailVerification();
        emit(const AuthStateNeedsVerification(
          isLoading: false,
          needCooldown: true,
        ));
      } catch (e) {
        emit(AuthStateRegistering(
          exception: e as Exception,
          isLoading: false,
        ));
      }
    });

    on<AuthEventSendEmailVerification>((event, emit) async {
      await authService.sendEmailVerification();
      emit(state);
    });

    on<AuthEventForgotPassword>((event, emit) async {
      final email = event.email;
      if (email == null) {
        emit(const AuthStateForgotPassword(
          exception: null,
          isLoading: false,
          emailSent: false,
        ));
        return; // user just opened the forgot password screen
      }

      // user wants to actually send a forgot-password email;
      emit(const AuthStateForgotPassword(
        isLoading: true,
        exception: null,
        emailSent: false,
      ));

      bool didSendEmail;
      Exception? exception;
      try {
        await authService.sendPasswordReset(toEmail: email);
        didSendEmail = true;
        exception = null;
      } on Exception catch (e) {
        didSendEmail = false;
        exception = e;
      }

      emit(AuthStateForgotPassword(
        exception: exception,
        isLoading: false,
        emailSent: didSendEmail,
      ));
    });
  }
}
