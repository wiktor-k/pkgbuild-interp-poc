# Maintainer: kpcyrd <kpcyrd[at]archlinux[dot]org>
# Contributor: Thomas Scholtes <geigerzaehler@axiom.fm>

pkgname=radicle-desktop
pkgver=0.17.0
pkgrel=1
pkgdesc='Radicle desktop app'
url='https://radicle.network/desktop'
arch=('x86_64')
license=('GPL-3.0-only')
depends=(
  'cairo' 'libcairo.so'
  'dbus' 'libdbus-1.so'
  'gdk-pixbuf2' 'libgdk_pixbuf-2.0.so'
  'glib2' 'libgio-2.0.so' 'libglib-2.0.so' 'libgobject-2.0.so'
  'glibc'
  'gtk3' 'libgdk-3.so' 'libgtk-3.so'
  'hicolor-icon-theme'
  'libgcc' 'libgcc_s.so'
  'libsoup3' 'libsoup-3.0.so'
  'radicle'
  'webkit2gtk-4.1' 'libjavascriptcoregtk-4.1.so' 'libwebkit2gtk-4.1.so'
  'zlib' 'libz.so'
)
makedepends=(
  'cargo'
  'cargo-tauri'
  'git'
  'npm'
  'pango'
)
source=(
  "radicle-desktop::git+https://seed.radicle.dev/z4D5UCArafTzTQpDZNQRuqswh3ury.git#tag=releases/${pkgver}"
)
sha256sums=('f8d952d34dc630cf8f16d2cf1e1159bc6a866447678455d2f46d97f5b9bb29cb')
b2sums=('2fdac95f230fd0aab1e509ab6e71416233f0e335112360901ba0d60e0e967b56a477c0781400b44d46c464c7c8e93f555788ed4688d5f82052d8d9c7fd0d35ef')

prepare() {
  cd "${pkgname}"
  cargo fetch --locked --target host-tuple
  npm ci
}

build() {
  cd "${pkgname}"
  export CFLAGS+=" -ffat-lto-objects"
  cargo tauri build -b deb --ci
}

package() {
  mv -v "${pkgname}"/target/release/bundle/deb/radicle-desktop_*/data/usr "${pkgdir}"
}

# vim: ts=2 sw=2 et:
