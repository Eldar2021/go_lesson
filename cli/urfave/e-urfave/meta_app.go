package main

import (
	"log"
	"os"
)

func CreateMetaApp(name string) {
	if libExist, err := CheckLibFolder(); err == nil && libExist {
		createApp()
		createAppView()
		createAppLogic()
		createAppRouter()
	}
}

func createApp() {
	appCode := []byte(
		`// exporting files
export 'logic/app_cubit.dart';
export 'router/app_router.dart';
export 'view/app_view.dart';`)
	if err := os.Mkdir("./lib/app", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./lib/app/app.dart", appCode, 0644); err != nil {
		log.Fatal(err)
	}
}

func createAppView() {
	appViewCode := []byte(
		`import 'package:flutter/material.dart';
import '../app.dart';		

class MyApp extends StatelessWidget {
  const MyApp({super.key});
  
  @override
  Widget build(BuildContext context) {
	return MaterialApp(
	  title: 'Flutter App',
	  theme: ThemeData.dark(useMaterial3: true),
	  onGenerateRoute: AppRouter.onGenerateRoute,
	);
  }
}
`)

	if err := os.Mkdir("./lib/app/view", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./lib/app/view/app_view.dart", appViewCode, 0644); err != nil {
		log.Fatal(err)
	}
}

func createAppLogic() {
	appCubitCode := []byte(
		`import 'package:equatable/equatable.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

part 'app_state.dart';

class AppCubit extends Cubit<AppState> {
  AppCubit() : super(const AppState());
}
`)

	appStateCode := []byte(`part of 'app_cubit.dart';

class AppState extends Equatable{
	const AppState();

    @override
    List<Object> get props => [];
}
`)

	if err := os.Mkdir("./lib/app/logic", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./lib/app/logic/app_cubit.dart", appCubitCode, 0644); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./lib/app/logic/app_state.dart", appStateCode, 0644); err != nil {
		log.Fatal(err)
	}
}

func createAppRouter() {
	appRouterCode := []byte(
		`import 'package:flutter/material.dart';

class AppRouter {
  const AppRouter();

  static const String main = '/';

  static Route<void> onGenerateRoute(RouteSettings settings) {
    return switch (settings.name) {
      main => MaterialPageRoute(builder: (_) => const Scaffold()),
      _ => throw Exception('no builder specified for route named: [${settings.name}]'),
    };
  }
}
`)

	if err := os.Mkdir("./lib/app/router", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./lib/app/router/app_router.dart", appRouterCode, 0644); err != nil {
		log.Fatal(err)
	}
}
