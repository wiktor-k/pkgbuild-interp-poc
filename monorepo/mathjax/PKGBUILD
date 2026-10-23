# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Contributor: Kevin Dodd <jesin00@gmail.com>
# Contributor: Tianjiao Yin <ytj000@gmail.com>

pkgname=mathjax
pkgver=4.1.3
pkgrel=1
pkgdesc='An open source JavaScript display engine for mathematics that works in all modern browsers'
url='https://www.mathjax.org/'
arch=(any)
license=(Apache-2.0)
makedepends=(git)
source=(git+https://github.com/mathjax/MathJax#tag=$pkgver)
sha256sums=('168e090295bea0741f7a3f7ec271a2fdede95150732a32e9e3685cde9bf5d65f')

package() {
  cd MathJax
  mkdir -p "$pkgdir"/usr/share/mathjax
  cp -ar * "$pkgdir"/usr/share/mathjax
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname
}
