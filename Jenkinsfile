node{
    ws("$GOPATH/src/github.com/msalcantara/godog-jenkins") {
        stage('checkout repository'){
            git branch: 'master',
            url: 'https://github.com/msalcantara/godog-jenkins'
        }
        
        stage('pull dependencies'){
            sh 'go get github.com/DATA-DOG/godog/cmd/godog'
        }
        
        stage('run tests'){
            sh 'build.sh'
        }
        
        stage('cucumber reports'){
                step([$class: 'CucumberReportPublisher',
                jenkinsBasePath: 'http://localhost:8080/job/github-api',
                fileIncludePattern: '*.json',
                fileExcludePattern: '',
                jsonReportDirectory: 'cucumber/',
                ignoreFailedTests: false,
                missingFails: false,
                pendingFails: false,
                skippedFails: false,
                undefinedFails: false,
                parallelTesting: false])
        }
        
    }
}