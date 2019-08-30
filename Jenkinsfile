pipeline {
    agent {
        docker { image 'golang:1.12.9-stretch'}
    }
    environment {
            GO111MODULE = 'on'
    }
    stages {
        stage('Pre Test'){
            echo 'Pulling Dependencies'
            sh 'go version'
            sh 'go get -u github.com/golang/dep/cmd/dep'
            sh 'cd $GOPATH/src/cmd/project && dep ensure'
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