# Maintainer: Pierre Schmitz <pierre@archlinux.de>

pkgbase=php-apcu
pkgname=('php-apcu' 'php-legacy-apcu')
pkgver=5.1.28
pkgrel=1
arch=('x86_64')
pkgdesc='A userland caching module for PHP'
url='https://pecl.php.net/package/APCu'
makedepends=('php' 'php-legacy')
license=('PHP')
source=("https://pecl.php.net/get/apcu-$pkgver.tgz"
        'apcu.ini')
sha256sums=('ca9c1820810a168786f8048a4c3f8c9e3fd941407ad1553259fb2e30b5f057bf'
            '18b2d904848b185bdc7c0c6a5f7c82ec809e9ed3f137cd6d3420160f4756630f')

prepare() {
	cd "$srcdir/apcu-$pkgver"

	cp -a "$srcdir/apcu-$pkgver" "$srcdir/apcu-legacy-$pkgver"
}

build() {
	cd "$srcdir/apcu-$pkgver"
	phpize
	./configure --prefix=/usr
	make

	cd "$srcdir/apcu-legacy-$pkgver"
	phpize-legacy
	./configure --prefix=/usr
	make
}

check() {
	export REPORT_EXIT_STATUS=1
	export NO_INTERACTION=1
	export SKIP_ONLINE_TESTS=1
	export SKIP_SLOW_TESTS=1

	for dir in "$srcdir/apcu"{,-legacy}-$pkgver; do
		pushd "$dir"
		make test
		popd
	done
}

package_php-apcu() {
	depends=('php')
	backup=('etc/php/conf.d/apcu.ini')

	cd "$srcdir/apcu-$pkgver"

	make INSTALL_ROOT="$pkgdir" install

	install -D -m644 "$srcdir/apcu.ini" "$pkgdir/etc/php/conf.d/apcu.ini"
	install -D -m644 apc.php "$pkgdir/usr/share/webapps/php-apcu/apc.php"
}

package_php-legacy-apcu() {
	depends=('php-legacy')
	backup=('etc/php-legacy/conf.d/apcu.ini')

	cd "$srcdir/apcu-legacy-$pkgver"

	make INSTALL_ROOT="$pkgdir" install

	install -D -m644 "$srcdir/apcu.ini" "$pkgdir/etc/php-legacy/conf.d/apcu.ini"
	install -D -m644 apc.php "$pkgdir/usr/share/webapps/php-legacy-apcu/apc.php"
}
