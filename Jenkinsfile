pipeline {
    agent any
    tools { go '1.21.1' }
    stages {    
        stage('Pre test') {
            steps {
                 sh 'go version'
            }
        } // Тут була пропущена фігурна дужка
        
        stage('Build') {
            steps {
                sh "ls"
                echo 'Compiling and building'
                sh 'go run ./cmd/main.go' // замість 'go run'
            }
        }
    }    
}
