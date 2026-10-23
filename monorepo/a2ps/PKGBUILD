# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: Eric Bélanger <eric@archlinux.org>

pkgname=a2ps
pkgver=4.15.8
pkgrel=1
pkgdesc='An Any to PostScript filter'
arch=(x86_64)
url="https://www.gnu.org/software/$pkgname"
license=(GPL-3.0-or-later)
depends=(imagemagick
         perl
         psutils)
makedepends=(gperf)
backup=('etc/a2ps/a2ps.cfg'
        'etc/a2ps/a2ps-site.cfg')
_archive="$pkgname-$pkgver"
source=("https://ftp.gnu.org/gnu/$pkgname/$_archive.tar.gz"{,.sig}
        "$pkgname-4.13c-emacs.patch"
        "$pkgname-adhere-source-date-epoch.patch")
sha256sums=('8d13915a36ebbfa8e7b236b350cc81adc714acb217a18e8d8c60747c0ad353f9'
            'SKIP'
            '943739a788db47f10942477754f42c8c9b22cdaad4463bea4103c31a330de76a'
            '305ba7b9a44e7d6480a23098efb27aab5daedde71defcfe2f1b01b70d025e3db')
validpgpkeys=('24093F016FFE8602EF449BB84C8EF3DA3FD37230') # Reuben Thomas <rrt@sc3d.org>

prepare() {
  cd "$_archive"
  patch -p1 -i "../$pkgname-4.13c-emacs.patch"
  patch -p1 -i "../$pkgname-adhere-source-date-epoch.patch"
}

build() {
	cd "$_archive"
	libtoolize --force --copy
	autoreconf --force --install -I m4
	LIBS+="-lm" ./configure \
		--prefix=/usr \
		--sysconfdir=/etc/a2ps \
		--includedir=/usr/include \
		--enable-shared \
		--enable-nls
	make
}

package() {
	cd "$_archive"
	make DESTDIR="$pkgdir" install
}
