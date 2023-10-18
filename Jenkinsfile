pipeline {
    agent any

    stages {    
        stage('Setup Go') {
            steps {
                script {
                    def GO_VERSION = '1.21.1'
                    sh """
                        wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz
                        sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
                        echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.bashrc
                        source $HOME/.bashrc
                    """
                }
                sh 'go version'
            }
        
        stage('Build') {
            steps {
                sh "ls"
                echo 'Compiling and building'
                sh 'go run ./cmd/main.go'
            }
        }


        }
        
    }
