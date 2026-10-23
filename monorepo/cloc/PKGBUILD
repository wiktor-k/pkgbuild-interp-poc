# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Evangelos Foutras <foutrelis@archlinux.org>
# Contributor: Diogo Ferreira <diogo@underdev.org>

pkgname=cloc
pkgver=2.10
pkgrel=1
pkgdesc="Count lines of code"
arch=('any')
url="https://github.com/AlDanial/cloc"
license=('GPL-2.0-or-later')
depends=(
  'perl'
  'perl-algorithm-diff'
  'perl-regexp-common'
  'perl-parallel-forkmanager'
)
makedepends=('git')
checkdepends=('unzip')
source=("git+$url.git#tag=v$pkgver")
b2sums=('08392540a3ea4e5cd463cf102a675dd7bd47be4d81e59529c6b7fcced57bcbc8b66300195016d52e33cd25168dd85c09a493a17290b2ae21b625b57c8478f10f')

check() {
  cd $pkgname/Unix
  make test-perl test-pod
  ./t/00_C.t
  # ./t/02_git.t # Clones git repo / requires network access
  ./t/01_opts.t
}

package() {
  cd $pkgname/Unix
  make DESTDIR="$pkgdir" install
}
