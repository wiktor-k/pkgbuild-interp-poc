# Maintainer: Morten Linderud <foxboron@archlinux.org>
# Maintainer: Frederik Schwan <freswa at archlinux dot org>

pkgname=slirp4netns
pkgdesc='User-mode networking for unprivileged network namespaces'
pkgver=1.3.6
pkgrel=1
arch=(x86_64)
url="https://github.com/rootless-containers/slirp4netns"
license=(GPL-2.0-or-later)
depends=(glibc glib2 libcap libseccomp libslirp)
makedepends=(git)
source=("git+$url#tag=v$pkgver?signed")
sha256sums=('ab9f5919ac371b410b6d325bd4cf6a635eda077faacee87c706d183f0b4284dd')
validpgpkeys=('C020EA876CE4E06C7AB95AEF49524C6F9F638F1A') # Akihiro Suda <akihiro.suda.cz@hco.ntt.co.jp>

prepare() {
  cd $pkgname
  autoreconf -fi
}

build() {
  cd $pkgname
  ./configure --prefix=/usr
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
