# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-tpm2-pytss
_name=${pkgname#python-}
pkgver=3.0.0
pkgrel=1
pkgdesc='Python bindings for tpm2-tss'
arch=(x86_64)
url='https://github.com/tpm2-software/tpm2-pytss'
license=(BSD-2-Clause)
depends=(
  glibc
  python
  python-asn1crypto
  python-cffi
  python-cryptography
  python-packaging
  python-yaml
)
makedepends=(
  git
  python-build
  python-installer
  python-pkgconfig
  python-setuptools
  python-setuptools-scm
  python-wheel
  tpm2-tss
)
checkdepends=(
  python-pytest
  swtpm
  tpm2-tools
)
source=(
  "git+$url.git#tag=$pkgver"
)
sha512sums=('3f433c96bfcd0c3eaf8f5b1659b41d4bcacd46fecc87961580f02e084c2c22d833c1337b4d8825eefa63cb9ad2adec167aa7dd62c4746948c6d59ffc77d5ce55')

build() {
  cd $_name
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
  )

  cd $_name
  local _python_version=$(python -c 'import sys; print("".join(map(str, sys.version_info[:2])))')
  PYTHONPATH="$PWD/build/lib.linux-$CARCH-cpython-$_python_version" pytest "${pytest_options[@]}"
}

package() {
  depends+=(
    tpm2-tss libtss2-esys.so libtss2-fapi.so libtss2-mu.so libtss2-rc.so libtss2-tctildr.so
  )

  cd $_name
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname"
}
