pipeline {
    agent any
    tools { go '1.21.1' }
    stages {    
        stage('Pre test') {
            steps {
                 sh 'go version'
            }
        } 

         stage('Load modules') {
            steps {
                script {
                   sh "go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
                }
            }
        }

        stage('Lint') {
            steps {
                script {
                  sh '$(go env GOPATH)/bin/golangci-lint run --disable-all -E errcheck'
                }
            }
        }
        
        stage('Migrate DB'){
            steps{
                sh "migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up"
            }
        }
              
        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go build ./cmd/main.go'
            }
        }
    }    
}
