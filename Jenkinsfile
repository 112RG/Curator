node {
  def app
  def dockerfile = 'Dockerfile'
  def branch = ''
  stage('Clone repo'){
    checkout scm
    script{
      env.GIT_COMMIT = sh(script: "git describe --always", returnStdout: true)
      env.BUILD_DATE = sh(script: "date -u +%Y-%m-%dT%H:%M", returnStdout: true)
    }
  }
  stage('Build image'){
    echo "Git rev ${GIT_COMMIT}"
    app = docker.build("artifactory.othala.xyz/orilla-docker-local/curator:latest_${env.BRANCH_NAME}", ". --build-arg GIT_COMMIT='${GIT_COMMIT}' --build-arg BUILD_DATE='${BUILD_DATE}'")
  }

  stage('Push image'){
    docker.withRegistry('https://artifactory.othala.xyz', 'OTHALA_REGISTRY'){
          app.push()
    }
  }
}