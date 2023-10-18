pipeline {
    agent any
    tools { go '1.21.1' }
    stages {    
        stage('Pre test') {
            steps {
                 sh 'go version'
            }
        } 
        stage('Load linter') {
                    steps {
                script {
sh 'curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.1'
                }
            }
        }
         stage('Load modules') {
            steps {
                script {
                    sh "go mod download"
                }
            }
        }

        stage('Lint') {
            steps {
                script {
                   sh "$GOPATH/bin/golangci-lint run -v"
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
                sh "ls"
                echo 'Compiling and building'
                sh 'go build ./cmd/main.go'
            }
        }

        stage('Run') {
            steps {
                sh './main'
            }
        }
    }    
}
