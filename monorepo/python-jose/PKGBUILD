# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-jose
pkgver=3.5.0
pkgrel=2
pkgdesc="A JOSE implementation in Python"
arch=('any')
license=('MIT')
url="https://github.com/mpdavis/python-jose"
depends=('python-ecdsa' 'python-pyasn1' 'python-rsa')
optdepends=('python-cryptography: one option for better performance'
            'python-pycryptodome: one option for better performance')
makedepends=('git' 'python-setuptools' 'python-build' 'python-installer' 'python-wheel')
checkdepends=('python-cryptography' 'python-pycryptodome' 'python-pytest')
source=("git+https://github.com/mpdavis/python-jose.git#tag=$pkgver"
	"fix-test_incorrect_public_key_hmac_signing.patch")
sha512sums=('1cb4d1768aa39e556312b4952a09ea51287a920cf941f477a56679c55c3ec06f7e54787e4579915d16a0c3e6744c220e0f2c0495cf5a53c17285ffde7ad964cb'
            '20f89f1c4d92e98eeb60114d60cdc5d51abe33de84220422323983b56a815ec2bcc0eb5cb70777972781b15978a78a388374d2a45b4fcfdd208e3831fcf3d229')

prepare() {
  cd python-jose
  # https://github.com/mpdavis/python-jose/pull/338
  git cherry-pick -n 574e3a340b53e70e2c42212d18bb2fe1c35d7d68
  # Fix https://github.com/mpdavis/python-jose/pull/393
  patch -Np1 -i ${srcdir}/fix-test_incorrect_public_key_hmac_signing.patch
}

build() {
  cd python-jose
  python -m build --wheel --no-isolation --skip-dependency-check
}

check() {
  cd python-jose
  pytest
}

package() {
  cd python-jose
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
