# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Yangtse <yangtsesu@gmail.com>
# Contributor: WU Jun <quark at lihdd dot net>

pkgname=libpinyin
pkgver=2.10.3
_modelver=20
pkgrel=2
pkgdesc="Library to deal with pinyin"
arch=('x86_64')
url="https://github.com/libpinyin/libpinyin"
license=('GPL-3.0-or-later')
depends=('kyotocabinet' 'glib2')
makedepends=('git')
source=("git+https://github.com/libpinyin/libpinyin.git#tag=$pkgver"
        "$pkgname-model$_modelver.text.tar.gz::https://downloads.sourceforge.net/libpinyin/model$_modelver.text.tar.gz")
noextract=("$pkgname-model$_modelver.text.tar.gz")
sha512sums=('7402fcdf759f608fdb1ce8cb397f95e9872b2ef6ee09a5d70a4950c7a0051bd9f7095cbdbd12e97dff32ada3a15f3f5df73c1cf0a446163dfaea1d34a9b8ccbf'
            'ed4d0607ad35e0e7ea424670539ddcd81a2b03c1da914b9c00cb748cf065f29471502d40b9a189852001da1fb9178c3bcc4675d7efebea5d081d78bfeee9b5d6')

prepare() {
  cd libpinyin

  ln -s "$srcdir"/$pkgname-model$_modelver.text.tar.gz data/model$_modelver.text.tar.gz
  sed -i "/wget.*model$_modelver\.text\.tar\.gz/ d" data/Makefile.am
}

build() {
  cd libpinyin
  ./autogen.sh --prefix=/usr --with-dbm=KyotoCabinet --enable-libzhuyin
  make
}

package() {
  cd libpinyin
  make DESTDIR="$pkgdir" install
}
