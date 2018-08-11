var gulp = require("gulp"),
    browserify = require("browserify"),
    source = require('vinyl-source-stream'),
    watchify = require("watchify"),
    tsify = require("tsify"),
    gutil = require("gulp-util"),
    sass = require("gulp-sass"),
    browserSync = require("browser-sync");

var paths = {
    pages: ['pages/**/*.html'],
    scss: ['scss/**/*.scss']
};

gulp.task('browser-sync', function() { // Создаем таск browser-sync
    browserSync({ // Выполняем browserSync
        server: { // Определяем параметры сервера
            baseDir: '../dist/' // Директория для сервера - app
        },
        notify: false // Отключаем уведомления
    });
});

var watchedBrowserify = watchify(browserify({
    basedir: '.',
    debug: true,
    entries: ['scripts/ts/main.ts'],
    cache: {},
    packageCache: {}
}).plugin(tsify));

gulp.task("copy-html", function() {
    return gulp
        .src(paths.pages)
        .pipe(gulp.dest("../dist"));
});

gulp.task("scss", function() {
    return gulp
        .src(paths.scss)
        .pipe(sass().on('error', sass.logError))
        .pipe(gulp.dest("../dist/css"));
});

gulp.task('watch', ['browser-sync', 'copy-html'], function() {
    gulp.watch('../dist/**/*', browserSync.reload); // Наблюдение за JS файлами в папке js
    gulp.watch('pages/**/*.html', ['copy-html', browserSync.reload]);
    gulp.watch(paths.scss, ['scss', browserSync.reload]);
});

function bundle() {
    return watchedBrowserify
        .bundle()
        .pipe(source('bundle.js'))
        .pipe(gulp.dest("../dist/js"));
}

gulp.task("default", ["watch"], bundle); ///*["copy-html"],*/

watchedBrowserify.on("update", bundle);
watchedBrowserify.on("log", gutil.log);