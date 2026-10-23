# Maintainer: David Runge <dvzrv@archlinux.org>

_name=redis
_upstream=phpredis
pkgbase=php-redis
pkgname=(
  php-redis
  php-legacy-redis
)
pkgver=6.3.0
pkgrel=2
pkgdesc="An API for communicating with the Redis key-value store"
arch=(x86_64)
url="https://github.com/phpredis/phpredis/"
license=(PHP-3.01)
depends=(glibc)
makedepends=(
  liblzf
  zstd
  php-igbinary
  php-legacy-igbinary
)
checkdepends=(
  lsof
  valkey
)
optdepends=('valkey: use a local redis instance')
source=($pkgname-$pkgver.tar.gz::$url/archive/$pkgver.tar.gz)
sha512sums=('44050095de7815547ca6f0ae263ea3a71eb7d1db9164cbc1df496ac25698546ae8fd3b9bd98863039b8ea5116f36f2d0fb3063370ab427b15472628939d8c7c5')
b2sums=('c35bb8b5d487c9350e4e988de48ba25dfdaf094d66fd390fc09959635cc78ae70860a3e3d617acea24f33e1f72886d39465df362eea72b0f322e058926333c41')

prepare() {
  mv -v $_upstream-$pkgver $pkgbase-$pkgver
  # tempfile is non-standard, Debian only
  sed -e 's/tempfile/mktemp/g' -i  $pkgname-$pkgver/tests/mkring.sh
  # the kill after shutdown of redis makes it exit with status code 1
  sed -e '/kill -9/d' -i $pkgname-$pkgver/tests/mkring.sh
  # disable the extension by default
  echo -e "; this extension requires igbinary to be activated as well\n;extension=${_name}" > "$pkgname-$pkgver/${_name}.ini"

  cp -av $pkgbase-$pkgver "${pkgname[1]}-$pkgver"

  (
    cd $pkgbase-$pkgver
    phpize
  )
  (
    cd "${pkgname[1]}-$pkgver"
    phpize-legacy
  )
}

build() {
  local configure_options=(
    --prefix=/usr
    --enable-redis-igbinary
    --enable-redis-zstd
    --enable-redis-lzf
    --with-liblzf=/usr/lib/
    --with-libzstd=/usr/lib/
  )

  (
    cd $pkgname-$pkgver
    ./configure "${configure_options[@]}"
    make
  )
  (
    cd "${pkgname[1]}-$pkgver"
    ./configure "${configure_options[@]}"
    make
  )
}

check() {
  # tests are partly broken:
  # https://github.com/phpredis/phpredis/issues/1593
  (
    export TEST_PHP_EXECUTABLE=/usr/bin/php
    export TEST_PHP_ARGS="-d extension=igbinary -d extension=$srcdir/$pkgbase-$pkgver/modules/redis.so"
    cd "$pkgname-$pkgver"
    tests/mkring.sh start
    $TEST_PHP_EXECUTABLE $TEST_PHP_ARGS tests/TestRedis.php --class Redis
    $TEST_PHP_EXECUTABLE $TEST_PHP_ARGS tests/TestRedis.php --class RedisArray
    tests/mkring.sh stop
  )
  (
    export TEST_PHP_EXECUTABLE=/usr/bin/php-legacy
    export TEST_PHP_ARGS="-d extension=igbinary -d extension=$srcdir/${pkgname[1]}-$pkgver/modules/redis.so"
    cd "${pkgname[1]}-$pkgver"
    tests/mkring.sh start
    $TEST_PHP_EXECUTABLE $TEST_PHP_ARGS tests/TestRedis.php --class Redis
    $TEST_PHP_EXECUTABLE $TEST_PHP_ARGS tests/TestRedis.php --class RedisArray
    tests/mkring.sh stop
  )
}

package_php-redis() {
  depends+=(liblzf.so libzstd.so php-igbinary)
  backup=(etc/php/conf.d/$_name.ini)

  cd $pkgname-$pkgver
  make INSTALL_ROOT="$pkgdir/" install
  install -vDm 644 $_name.ini -t "$pkgdir/etc/php/conf.d/"
  install -vDm 644 {{CHANGELOG,README,arrays,cluster}.md,CREDITS} -t "$pkgdir/usr/share/doc/$pkgname/"
}

package_php-legacy-redis() {
  depends+=(liblzf.so libzstd.so php-legacy-igbinary)
  backup=(etc/php-legacy/conf.d/$_name.ini)

  cd "$pkgname-$pkgver"
  make INSTALL_ROOT="$pkgdir/" install
  install -vDm 644 $_name.ini -t "$pkgdir/etc/php-legacy/conf.d/"
  install -vDm 644 {{CHANGELOG,README,arrays,cluster}.md,CREDITS} -t "$pkgdir/usr/share/doc/$pkgname/"
}
