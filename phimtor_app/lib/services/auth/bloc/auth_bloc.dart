import 'package:bloc/bloc.dart';
import 'package:phimtor_app/services/auth/auth_provider.dart';
import 'package:phimtor_app/services/auth/bloc/auth_event.dart';
import 'package:phimtor_app/services/auth/bloc/auth_state.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  static AuthState getInitialState(AuthProvider provider) {
    final user = provider.currentUser;
    return user == null
        ? const AuthStateLoggedOut(exception: null, isLoading: false)
        : user.emailVerified == false
            ? const AuthStateNeedsVerification(isLoading: false)
            : AuthStateLoggedIn(user: user, isLoading: false);
  }

  AuthBloc(AuthProvider provider) : super(getInitialState(provider)) {
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
        final user = await provider.logIn(email: email, password: password);
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
        await provider.logOut();
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
        await provider.createUser(email: email, password: password);
        await provider.sendEmailVerification();
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
      await provider.sendEmailVerification();
      emit(state);
    });
  }
}
