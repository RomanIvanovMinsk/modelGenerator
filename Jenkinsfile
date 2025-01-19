pipeline {
    agent any
    environment {
        SSH_CREDENTIALS_ID = 'bc013f38-40d9-4731-8ed1-23c56055cc0f' // Replace with your SSH credential ID from you jenkin dashboard
    }
    stages {
        stage('Build') {
            steps {
                // Build the Go application
                sh 'go build -o modelGenerator' 
        }
        stage('Test') {
            steps {
                // Run Go unit tests
                sh 'go test -v'
            }
        }
      }
    }
}