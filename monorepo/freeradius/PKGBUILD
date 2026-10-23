# Maintainer: Sergej Pupykin <pupykin.s+arch@gmail.com>
# Contributor: Jason R Begley (jayray@digitalgoat.com>
# Contributor: Daniel Micay <danielmicay@gmail.com>

pkgname=freeradius
pkgver=3.2.10
pkgrel=4
pkgdesc='The premier open source RADIUS server'
arch=('x86_64')
url='https://freeradius.org/'
license=('GPL-2.0-or-later')
depends=('krb5' 'net-snmp' 'postgresql-libs' 'mariadb-libs' 'talloc'
         'libpcap' 'libxcrypt' 'libcrypt.so' 'smbclient')
makedepends=('curl' 'freetds' 'hiredis' 'json-c' 'libmemcached' 'perl'
             'python' 'samba' 'sqlite' 'unixodbc' 'yubico-c-client')
optdepends=('curl: for REST'
            'freetds: for Sybase and MS SQL'
            'hiredis: for redis support'
            'json-c: rlm_rest module'
            'libmemcached: for memcached'
            'perl: for Perl'
            'python: for Python'
            'sqlite: for sqlite'
            'unixodbc: for ODBC'
            'yubico-c-client: for yubicloud')
options=('!makeflags')
validpgpkeys=('BF2384EC6938B9744B03E2A620E37C25995B4F85') # FreeRADIUS - Package Signing [...] <packages@freeradius.org>
source=("https://github.com/FreeRADIUS/freeradius-server/releases/download/release_${pkgver//./_}/freeradius-server-${pkgver}.tar.bz2"{,.sig}
        'freeradius.service'
        'freeradius.sysusers'
        'freeradius.tmpfiles')
sha256sums=('e30ce4b071ad4b48e7eb0c34c66324a32c2548cfc067900941a90f017dc174e3'
            'SKIP'
            'ff9e31232122d2982fae5466969f45f9f46396c307754aa10b582e5cab19758f'
            '717e0cbf090d023dde5f8aba6850ce9b25da29aa28dd6b5b4f98535b5872f14f'
            'f536a9aa972e3e42a6b1a6d8ee17166eb721c7cba2e80f60473811497c7bd8bc')

build() {
  cd "$srcdir"/freeradius-server-$pkgver

  ./configure \
    --prefix=/usr \
    --sbindir=/usr/bin \
    --sysconfdir=/etc \
    --libdir=/usr/lib/freeradius \
    --localstatedir=/var \
    --with-unixodbc-dir=/usr/lib \
    --enable-heimdal-krb5 \
    --enable-reproducible-builds \
    --without-rlm_python \
    --with-regex=false \
    --with-rlm_python3 \
    --with-system-libltdl \
    --with-system-libtool \
    --with-udpfromto
  make PACKAGE='archlinux'
}

package() {
  cd "$srcdir"/freeradius-server-$pkgver

  make install R="$pkgdir" PACKAGE='archlinux'
  chmod o+r "$pkgdir"/etc/raddb/*
  mv "$pkgdir"/etc/raddb "$pkgdir"/etc/raddb.default
  rm -rf "$pkgdir"/var/run

  install -D -m0644 "$srcdir"/$pkgname.service "$pkgdir"/usr/lib/systemd/system/$pkgname.service
  install -D -m0644 "$srcdir"/$pkgname.sysusers "$pkgdir"/usr/lib/sysusers.d/$pkgname.conf
  install -D -m0644 "$srcdir"/$pkgname.tmpfiles "$pkgdir"/usr/lib/tmpfiles.d/$pkgname.conf
}
