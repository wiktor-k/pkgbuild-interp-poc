# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Gaetan Bisson <bisson@archlinux.org>
# Contributor: Angel Velasquez <angvp@archlinux.org>
# Contributor: Hugo Doria <hugo@archlinux.org>

pkgbase=nmap
pkgname=(
  nmap
  ndiff
  zenmap
)
pkgver=7.992
pkgrel=1
pkgdesc="Utility for network discovery and security auditing"
url='https://nmap.org/'
arch=(x86_64)
license=(LicenseRef-Nmap-Public-Source-License-Version-0.95)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-setuptools-gettext
  python-wheel
  # nmap deps:
  glibc
  libgcc
  libpcap
  libssh2
  libstdc++
  lua54
  openssl
  pcre2
  zlib
  # zenmap deps:
  gdk-pixbuf2
  glib2
  gtk3
  pango
  python
  python-cairo
  python-gobject
)
source=(
  "https://nmap.org/dist/$pkgname-$pkgver.tar.bz2"
  "https://nmap.org/dist/sigs/$pkgname-$pkgver.tar.bz2.asc"
)
sha512sums=('0c38610b78190b1865d5adb900e9bdc5813a29f6f5d47b768f4b1c9d0d36a26de476f7c7fc55b04e2421c6e6b149890ab0207ea1740786158bbf594a834fedc7'
            'SKIP')
b2sums=('378d50ee44657877f8a0ba4bb3c06f70790692da79e5e601c18dd73926ea7779ac6633cea126583d241205d0bfda890a8dea750cbbef388d4ffbd25d91e058a6'
        'SKIP')
validpgpkeys=('436D66AB9A798425FDA0E3F801AF9F036B9355D0') # Nmap Project Signing Key (http://www.insecure.org/)

prepare() {
  cd $pkgbase-$pkgver
  # Ensure we build de-vendored deps
  rm -r liblua libpcap libpcre macosx mswin32 libssh2 libz
  # Keep zenmap from installing ndiff; Arch packages it separately
  sed -e '/^ndiff = "ndiff:run_main"$/d' -e '/^py-modules = \["ndiff"\]$/d' -i zenmap/pyproject.toml

  autoreconf -fiv -I /usr/share/gettext/m4
}

build() {
  cd $pkgbase-$pkgver
  ./configure \
    --prefix=/usr \
    --with-libpcap=/usr \
    --with-libpcre=/usr \
    --with-zlib=/usr \
    --with-libssh2=/usr \
    --with-liblua=/usr \
    --without-ndiff \
    --without-zenmap
  make

  python -m build --wheel --no-isolation ndiff

  python -m build --wheel --no-isolation zenmap
}

check() {
  cd $pkgbase-$pkgver
  make check

  cd zenmap
  python -m unittest discover -p '*.py'

  cd ../ndiff
  python ndifftest.py
}

package_nmap() {
  depends=(
    glibc
    libgcc
    libpcap
    libssh2
    libstdc++
    lua54
    openssl
    pcre2
    zlib
  )

  cd $pkgbase-$pkgver
  make DESTDIR="$pkgdir" install
  install -vDm644 -t "$pkgdir/usr/share/doc/$pkgname" \
    README.md \
    docs/nmap.usage.txt
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" \
    LICENSE \
    docs/3rd-party-licenses.txt
}

package_ndiff() {
  pkgdesc="Compare two Nmap XML files and display a list of their differences"
  depends=(python)

  cd $pkgbase-$pkgver
  python -m installer --destdir="$pkgdir" ndiff/dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
  install -vDm644 -t "$pkgdir/usr/share/man/man1" ndiff/docs/ndiff.1
}

package_zenmap() {
  pkgdesc+=" (GUI)"
  depends=(
    gdk-pixbuf2
    glib2
    gtk3
    ndiff
    nmap
    pango
    python
    python-cairo
    python-gobject
  )

  cd $pkgbase-$pkgver
  python -m installer --destdir="$pkgdir" zenmap/dist/*.whl
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
  install -vDm644 -t "$pkgdir/usr/share/man/man1" docs/zenmap.1

  ln -vs /usr/bin/zenmap "$pkgdir/usr/bin/nmapfe"
  ln -vs /usr/bin/zenmap "$pkgdir/usr/bin/xnmap"

  install -vDm644 -t "$pkgdir/usr/share/applications" \
    zenmap/install_scripts/unix/zenmap-root.desktop \
    zenmap/install_scripts/unix/zenmap.desktop
  install -vDm644 -t "$pkgdir/usr/share/icons/hicolor/256x256/apps" \
    zenmap/zenmapCore/data/pixmaps/zenmap.png
}
