# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Contributor: Sergej Pupykin <pupykin.s+arch@gmail.com>
# Contributor: Sirius Bakke <sirius\at/nonline.org>

pkgname=pam_mount
pkgver=2.22
pkgrel=1
pkgdesc='A PAM module that can mount volumes for a user session'
arch=(x86_64)
url='http://pam-mount.sourceforge.net/'
license=(GPL-2.0-or-later
         LGPL-2.1-or-later)
depends=(cryptsetup
         glibc
         libhx
         libxml2
         openssl
         pam
         pcre2
         util-linux-libs)
backup=('etc/security/pam_mount.conf.xml')
options=(!emptydirs)
source=(https://inai.de/files/pam_mount/$pkgname-$pkgver.tar{.zst,.asc})
sha256sums=('26c1c73550ee2dcf65e30a988d65a3e0b72e516a817ac0051c46a03bd435ff82'
            'SKIP')
validpgpkeys=(2F8BFE96D14F8B560FBC1F1DA28BDFEEB5A2B2FF  # Jan Engelhardt <jengelh@inai.de>
              23686C10A45691BE7A425109D6388181F35A0938
              BCA0C5C309CAC569E74A921CF76EFE5D0C223A8F
              2BDA563E9DDE46AF71CC171DAEB6941A2F9C9944)

build() {
  cd $pkgname-$pkgver
  ./configure \
	--prefix=/usr \
	--with-ssbindir=/usr/bin \
	--sbindir=/usr/bin \
	--with-slibdir=/usr/lib \
	--sysconfdir=/etc \
	--localstatedir=/var
  sed -i -e 's/ -shared / -Wl,-O1,--as-needed\0/g' libtool # Fix overlinking
  make
}

package() {
  cd $pkgname-$pkgver
  make DESTDIR="$pkgdir" install
}
