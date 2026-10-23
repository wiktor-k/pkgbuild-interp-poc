# Maintainer: Jonathan Steel <jsteel at archlinux.org>
# Contributor: thorsten w. <p@thorsten-wissmann.de>

pkgname=herbstluftwm
pkgver=0.9.6
pkgrel=1
pkgdesc="Manual tiling window manager for X"
arch=('x86_64')
url="https://herbstluftwm.org"
license=('BSD-2-Clause')
depends=('libxinerama' 'libxrandr' 'libxft' 'libxext' 'libxfixes' 'libxrender' 'freetype2')
makedepends=('cmake' 'asciidoc')
optdepends=('xorg-xsetroot: to set wallpaper color in default autostart'
            'xterm: used as the terminal in default autostart'
            'dzen2: used in the default panel.sh'
            'dmenu: used in some example scripts')
source=(https://herbstluftwm.org/tarballs/$pkgname-$pkgver.tar.gz{,.sig})
sha256sums=('e38c9721cd5c835ec1d461ab55c9dabcacea8c42eb37aeb782be7b138c91c6eb'
            'SKIP')
validpgpkeys=('72B6C05CDFF309C6396167D464EF02821CAFF810') # Thorsten Wißmann <edu@thorsten-wissmann.de>

build() {
  cd $pkgname-$pkgver

  mkdir -p build; cd build

  cmake \
    -DBASHCOMPLETIONDIR=/usr/share/bash-completion/completions \
    -DZSHCOMPLETIONDIR=/usr/share/zsh/functions/Completion/X \
    -DCMAKE_INSTALL_PREFIX="/usr" ..

  make
}

package() {
  cd $pkgname-$pkgver/build

  make DESTDIR="$pkgdir" install

  install -d "$pkgdir"/usr/share/licenses/$pkgname/

  ln -s /usr/share/doc/$pkgname/LICENSE \
    "$pkgdir"/usr/share/licenses/$pkgname/LICENSE
}
