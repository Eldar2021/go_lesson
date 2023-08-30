package main

import (
	"log"
	"os"
)

func CreateFeature(name string) {
	if libExist, err := CheckLibFolder(); err == nil && libExist {
		createFetureFile(name)
		createFeatureView(name)
		createFeatureWidgets(name)
		createFeatureLogic(name)
	}
}

func createFetureFile(name string) {
	appCode := []byte(
		`// exporting files
export 'logic/` + name + `_cubit.dart';
export 'widgets/` + name + `_widgets.dart';
export 'view/` + name + `_view.dart';`)
	if err := os.Mkdir("./lib/"+name, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./lib/"+name+"/"+name+".dart", appCode, 0644); err != nil {
		log.Fatal(err)
	}
}

func createFeatureView(name string) {
	appViewCode := []byte(
		`import 'package:flutter/material.dart';

class ` + name + `View extends StatelessWidget {
  const ` + name + `View({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('` + name + ` Feature'),
      ),
      body: const Center(
        child: Text('` + name + ` Body'),
      ),
    );
  }
}
`)

	if err := os.Mkdir("./lib/"+name+"/view", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./lib/"+name+"/view/"+name+"_view.dart", appViewCode, 0644); err != nil {
		log.Fatal(err)
	}
}

func createFeatureWidgets(name string) {
	appRouterCode := []byte(`// Here are we will write our features widgets`)

	if err := os.Mkdir("./lib/"+name+"/widgets", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./lib/"+name+"/widgets/"+name+"_widgets.dart", appRouterCode, 0644); err != nil {
		log.Fatal(err)
	}
}

func createFeatureLogic(name string) {
	appCubitCode := []byte(
		`import 'package:equatable/equatable.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

part '` + name + `_state.dart';

class ` + name + `Cubit extends Cubit<` + name + `State> {
	` + name + `Cubit() : super(const ` + name + `State());
}
`)

	appStateCode := []byte(`part of '` + name + `_cubit.dart';

class ` + name + `State extends Equatable{
	const ` + name + `State();

    @override
    List<Object> get props => [];
}
`)

	if err := os.Mkdir("./lib/"+name+"/logic", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./lib/"+name+"/logic/"+name+"_cubit.dart", appCubitCode, 0644); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./lib/"+name+"/logic/"+name+"_state.dart", appStateCode, 0644); err != nil {
		log.Fatal(err)
	}
}
