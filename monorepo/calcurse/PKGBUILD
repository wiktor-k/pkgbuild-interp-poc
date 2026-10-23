# Maintainer: Lukas Fleischer <lfleischer@archlinux.org>
# Contributor: Ronald van Haren <ronald.archlinux.org>
# Contributor: wain <wain.at.gmx.net>

pkgname=calcurse
pkgver=4.8.2
pkgrel=1
pkgdesc='A text-based personal organizer'
arch=('x86_64')
url='https://calcurse.org/'
license=('BSD')
depends=('ncurses')
optdepends=(
  'python: for CalDAV support'
  'python-httplib2: for CalDAV support'
  'python-pyparsing: for CalDAV support'
  'python-oauth2client: for Google Calendar support'
)
source=("https://calcurse.org/files/${pkgname}-${pkgver}.tar.gz"{,.asc})
sha512sums=('9949eb32b79022c3c2e7c8cdf32de34cc370d7550e4211146666fadf2a00110b9c9f2e224b4ecd0caaf0b6437d73675b016f20caa123a3fef9d10ebe03c42460'
            'SKIP')
b2sums=('53951b8bfc94bab8847b1c1d819d2ac10e988725232d3ba84a37966c77b468cde86fa60b8fd3548657947792bdabcf32bf2db16a8d809c1b71a7b7664d59fbf4'
        'SKIP')
validpgpkeys=('2E36D8620221482FC45CB7F2A91764759326B440') # Lukas Fleischer <lfleischer@calcurse.org>

build() {
  cd "$pkgname-$pkgver"

  ./configure \
    --enable-docs \
    --without-asciidoc \
    --prefix=/usr \
    --mandir=/usr/share/man

  make
}

check() {
  cd "$pkgname-$pkgver"
  make check
}

package() {
  cd "$pkgname-$pkgver"

  make DESTDIR="$pkgdir" install

  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" COPYING
}
