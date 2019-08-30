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
                sh 'go version'
                sh 'go build .'
            }
        }
        stage('Test'){
            steps{
                sh 'go test -v'
            }
        }
    }
}