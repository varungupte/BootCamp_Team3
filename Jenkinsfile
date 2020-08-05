pipeline {
    agent any
    tools {
        go 'go1.14'
    }
    stages {
        stage('Build gin docker image') {
            steps {
                echo 'Stopping existing gin container'
                sh 'docker stop cont-go-gin'

                echo 'Remove existing gin image'
                sh 'docker rmi img-go-gin:1.1.1'

                echo 'Building new gin image'
                sh 'docker build -t img-go-gin:1.1.1 .'
            }
        }

        stage('Run gin docker container') {
            steps {
                echo 'Running new gin container'
                sh 'docker run --rm -d -p 7878:7878 --net=my_bridge --name=cont-go-gin img-go-gin:1.1.1'
            }
        }
    }
}
