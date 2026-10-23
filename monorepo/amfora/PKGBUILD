# Maintainer: Alexander F. Rødseth <xyproto@archlinux.org>
# Contributor: Paul Boyd <boyd.paul2@gmail.com>

pkgname=amfora
pkgver=1.11.0
pkgrel=1
pkgdesc='Terminal browser for the Gemini protocol'
arch=(x86_64)
url='https://github.com/makeworld-the-better-one/amfora'
license=(GPL-3.0-only)
depends=(glibc)
makedepends=(git go imagemagick)
source=("git+$url#tag=v$pkgver")
b2sums=('0d7c8a7c699090715805dc36f8d5c1bed78e3d321c24dec5d51d500267db4c29a9c81726a9b3627db154f4132419284ba3ddd81a443fe0ee49023df9766b683c')

prepare() {
  cd $pkgname
  convert logo.png -colorspace sRGB \( +clone -crop 50%x100%+0+0 \) -delete 0 -fuzz 10% -transparent white \
    -background none -resize 48x48 -gravity center -extent 48x48 -strip amfora.png
  echo Icon=amfora >> amfora.desktop
}

build() {
  cd $pkgname
  go build -buildmode=pie -mod=readonly -modcacherw -trimpath -ldflags \
    "-s -w -linkmode=external -extldflags \"$LDFLAGS\" -X main.version=$pkgver -X main.builtBy=Arch_Linux -X main.commit=${source#*=}"
}

package() {
  cd $pkgname
  install -Dm755 amfora "$pkgdir/usr/bin/amfora"
  install -Dm644 amfora.png "$pkgdir/usr/share/pixmaps/amfora.png"
  install -Dm644 amfora.desktop "$pkgdir/usr/share/applications/amfora.desktop"
}
