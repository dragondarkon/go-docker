pipeline {
    agent {
        docker { image 'golang:1.12.9-stretch'}
    }
    environment {
            GO111MODULE = 'on'
    }
    stages {
        stage('Pre Test'){
            steps{
               sh 'go version'
            }
        }
        stage('Build'){
            steps{
                sh 'go build .'
            }
        }
        stage('Test'){
            steps{
                sh 'go test -v'
            }
        }
        stage('Robot Test'){
            steps{
                sh 'robot automated_test/printHelloWorld.robot'
            }
        }
    }
}