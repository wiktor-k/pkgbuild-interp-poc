# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Contributor: Elmar Klausmeier <Elmar.Klausmeier@gmail.com>

pkgname=sundials
pkgver=7.9.0
pkgrel=1
pkgdesc='Suite of nonlinear differential/algebraic equation solvers'
arch=(x86_64)
url='https://computing.llnl.gov/projects/sundials'
license=(BSD-3-Clause)
depends=(glibc
         lapack
         libgomp
         openmpi
         suitesparse)
makedepends=(cmake
             gcc-fortran
             git
             python)
optdepends=('python-matplotlib: for some examples'
            'python-numpy: for some examples')
source=(git+https://github.com/LLNL/sundials#tag=v$pkgver)
sha256sums=('07235d2ba2c65518d5095390432601f668bb1d95bb3a3e517f2336dcf6c8b0e3')

build() {
  cmake -B build -S $pkgname \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -DBUILD_STATIC_LIBS=OFF \
    -DENABLE_MPI=ON \
    -DENABLE_PTHREAD=ON	\
    -DENABLE_OPENMP=ON \
    -DENABLE_KLU=ON \
    -DKLU_LIBRARY_DIR=/usr/lib \
    -DKLU_INCLUDE_DIR=/usr/include/suitesparse \
    -DENABLE_LAPACK=ON \
    -DSUNDIALS_EXAMPLES_INSTALL_PATH=/usr/share/sundials/examples
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build

  install -Dm644 $pkgname/LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname
}
