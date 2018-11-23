
var gulp = require("gulp");
var shell = require('gulp-shell')

//this compiles new binary with source change
gulp.task("install-binary", shell.task([
    'go build $curws/middleware.go'
]))

//second argument tells install-binary is a deapendency for restart-supervisor

gulp.task("restart-supervisor", ['install-binary'], shell.task([
    'sudo supervisorctl restart golang-project'
]))

gulp.task('watch', function() {
    //watch the source code for all changes
    gulp.watch("*", ['install-binary', 'restart-supervisor']);
});

gulp.task('default', ['watch']);
