# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: Christian Karl <christian dot karl at protonmail dot com>

pkgname=vimiv
pkgver=0.9.1
pkgrel=11
pkgdesc='Image viewer with vim-like keybindings'
url='https://github.com/karlch/vimiv'
arch=('x86_64')
license=('MIT')
depends=('python' 'python-gobject' 'gtk3' 'python-pillow' 'glibc')
makedepends=('python-build' 'python-installer' 'python-setuptools' 'python-wheel' 'libgexiv2')
optdepends=('libgexiv2: autorotation and EXIF support')
backup=('etc/vimiv/vimivrc'
        'etc/vimiv/keys.conf')
source=(${pkgname}-${pkgver}.tar.gz::https://github.com/karlch/vimiv/archive/v${pkgver}.tar.gz)
sha256sums=('7196341c41ad372f4d5d98bc96fba4aa55ad6e78d93afd617a62866bdaa6c087')
sha512sums=('b500e5b54450238b6471c1ea0c0831dc6b16238a99b9fcb531bc93bd0aa0b92fa192cbf2dd29f1679856f840dd6f6e4d27971403c85b71b6f333c525b77d9e69')

prepare() {
  # Prevent make install from running setup.py install 
  sed -i '/python3 setup\.py install/d' ${pkgname}-${pkgver}/Makefile
}

build() {
  cd ${pkgname}-${pkgver}
  python -m build --wheel --no-isolation
}

package() {
  cd ${pkgname}-${pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  make DESTDIR="${pkgdir}" install
  install -Dm 644 readme.md -t "${pkgdir}/usr/share/doc/${pkgname}"
  mv "$pkgdir"/usr/share/{appdata,metainfo}
}

# vim: ts=2 sw=2 et:
