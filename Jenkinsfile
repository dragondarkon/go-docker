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
        stage("Archive artifacts") {
            // Archive the binary files in Jenkins so we can retrieve them later should we need to audit them
            archiveArtifacts artifacts: 'binaries/**', fingerprint: true
        }
    }
}