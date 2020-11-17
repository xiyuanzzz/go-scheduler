#!/usr/bin/env groovy

pipeline {
  agent {
    label 'master'
  }
  options {
    timestamps()
  }
  environment {
    url = 'https://10.151.12.63/testmodule/autotest/'
    branch = 'master'
    credentials = '1c5d0843-8601-47e1-b082-0d12cd0ae323'
  }
  stages {
    stage('PREPARE PARAMS') {
      steps {
        echo 'params'
      }
    }
    stage('PREPARE TESTCASE') {
      steps {
        echo 'testcase'
      }
    }
    stage('EXECUTE TESTCASE) {
      steps {
        echo 'exec'
      }
    }
  }
}