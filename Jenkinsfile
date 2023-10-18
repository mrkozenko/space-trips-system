pipeline {
    agent any

    stages {    
        stage('Setup Go') {
            steps {
                 goInstall version: '1.21.1'
                 sh 'go version'
            }
        } // Тут була пропущена фігурна дужка
        
        stage('Build') {
            steps {
                sh "ls"
                echo 'Compiling and building'
                sh 'go build ./cmd/main.go' // замість 'go run'
            }
        }
    }    
}
