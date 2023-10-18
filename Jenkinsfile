pipeline {
    agent any

    stages {        
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'

            }
        }
        
        stage('Build') {
            steps {
                sh "ls"
                echo 'Compiling and building'
                sh 'go build /cmd'
            }
        }


        }
        
    }
