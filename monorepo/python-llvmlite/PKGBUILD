# Maintainer: Jakub Klinkovský <lahwaacz at archlinux dot org>
# Contributor: Konstantin Gizdov <arch at kge dot pw>

_name=llvmlite
pkgname=python-$_name
pkgver=0.50.0
pkgrel=1
pkgdesc="A lightweight LLVM Python binding for writing JIT compilers"
arch=(x86_64)
url="https://github.com/numba/llvmlite"
license=("BSD-2-Clause AND Apache-2.0 WITH LLVM-exception")
depends=(
    glibc
    libgcc
    libstdc++
    llvm22-libs
    python
)
makedepends=(
    cmake
    git
    llvm22
    python-build
    python-installer
    python-setuptools-scm
    python-wheel
)
checkdepends=(
    python-pytest
)
source=(git+https://github.com/numba/llvmlite.git#tag=v$pkgver)
b2sums=('ae47004289324855019b376c83df2085e79ea637e82133b96cf82d3dcb22c2007cda2f574ceffe835710af9c47eb1a2d6f1f63c616c28f78de48fdeeb90c4245')

build() {
    cd $_name
    export PATH="/usr/lib/llvm22/bin:$PATH"
    LLVMLITE_SHARED=ON python -m build --wheel --no-isolation
}

check() {
    cd $_name
    pytest -vv $_name/tests
}

package() {
    cd $_name

    python -m installer --destdir="$pkgdir" dist/*.whl

    install -vDm 644 LICENSE LICENSE.thirdparty -t "$pkgdir"/usr/share/licenses/$pkgname/
}
