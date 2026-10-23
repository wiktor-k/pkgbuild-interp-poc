# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Jonas Witschel <diabonas@archlinux.org>

pkgname=python-python-pkcs11
pkgver=0.10.1
pkgrel=1
pkgdesc='PKCS#11/Cryptoki support for Python'
arch=(x86_64)
url='https://github.com/pyauth/python-pkcs11'
license=(MIT)
depends=(
  glibc
  python
  python-asn1crypto
)
makedepends=(
  git
  cython
  python-build
  python-installer
  python-setuptools
  python-setuptools-scm
  python-wheel
)
checkdepends=(
  python-cryptography
  python-oscrypto
  python-parameterized
  python-pytest
  softhsm
)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('1111415c2a09656e2fa5eb092d5bac0ce1344d12471b30af45bb4e630163fbd29e112e9883d49cf99ee5203786746933b7896442d530f2a85dc44e08ea95f29f')
b2sums=('add74bbee9dec22520ebed341958451471799f0b2a1edd47616982b312045d8021215954251afe5b1a50f2eb6894d2b86762e7108f9d0468449253c1e8cdd628')

prepare() {
  cd "$pkgname"

  sed -n '/^Copyright/,$p' README.rst > LICENSE
}

build() {
  cd "$pkgname"

  SETUPTOOLS_SCM_PRETEND_VERSION=$pkgver python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  export PKCS11_MODULE=/usr/lib/libsofthsm2.so
  export PKCS11_TOKEN_LABEL=TEST
  export PKCS11_TOKEN_PIN=1234
  export PKCS11_TOKEN_SO_PIN=5678

  softhsm2-util \
    --init-token \
    --free \
    --label "${PKCS11_TOKEN_LABEL}" \
    --pin "${PKCS11_TOKEN_PIN}" \
    --so-pin "${PKCS11_TOKEN_SO_PIN}"

  local _python_version=$(python -c 'import sys; print("".join(map(str, sys.version_info[:2])))')
  export PYTHONPATH="$PWD/build/lib.linux-$CARCH-cpython-$_python_version"

  pytest --import-mode=append -k 'not test_aes_gcm_test_vector and not test_encrypt_gcm'
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
