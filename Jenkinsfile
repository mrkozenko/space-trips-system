pipeline {
    agent any
    tools { go '1.21.1' }
    stages {    
        stage('Pre test') {
            steps {
                 sh 'go version'
            }
        } // Тут була пропущена фігурна дужка

        stage('Migrate DB'){
steps{
    sh "migrate  -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up"
}
        }

        
        stage('Build') {
            steps {
                sh "ls"
                echo 'Compiling and building'
                sh 'go run ./cmd/main.go' // замість 'go run'
            }
        }
    }    
}
