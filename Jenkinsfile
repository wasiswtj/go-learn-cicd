// Based on:
// https://raw.githubusercontent.com/redhat-cop/container-pipelines/master/basic-spring-boot/Jenkinsfile

library identifier: "pipeline-library@v1.5",
retriever: modernSCM(
  [
    $class: "GitSCMSource",
    remote: "https://github.com/redhat-cop/pipeline-library.git"
  ]
)

// Application name
bcName = "go-learn-cicd"

pipeline {
    // Use the 'golang' Jenkins agent image which is provided with OpenShift 
    agent any
    tools {
        go 'go-1.13.8'
    }
    environment {
        GO111MODULE = 'on'
        CGO_ENABLED = 0 
    }
    stages {
        stage("Checkout") {
            steps {
                checkout scm
            }
        }
        stage("Unit Test") {
            steps {
                sh 'go test -v -coverprofile=coverage.out ./calculations/handler/ ./calculations/usecase/'
            }
        }
        stage("Docker Build") {
            steps {
                // This uploads your application's source code and performs a binary build in OpenShift
                // This is a step defined in the shared library (see the top for the URL)
                // (Or you could invoke this step using 'oc' commands!)
                binaryBuild(buildConfigName: bcName, buildFromPath: ".")
            }
        }
    }
}