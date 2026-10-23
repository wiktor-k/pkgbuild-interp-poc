# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: Josip Ponjavic <josipponjavic at gmail dot com>

pkgname=python-pycryptodome
pkgver=4.1.0
pkgrel=1
license=('BSD-2-Clause' 'Unlicense' 'MIT')
arch=('x86_64')
pkgdesc="Collection of cryptographic algorithms and protocols, implemented for use from Python 3."
url='https://www.pycryptodome.org/'
depends=('python' 'glibc')
makedepends=('git' 'python-'{build,installer,wheel} 'python-setuptools')
checkdepends=('python-pytest')
conflicts=('python-crypto')
provides=('python-crypto')
replaces=('python-crypto')
source=("git+https://github.com/Legrandin/pycryptodome.git#tag=v$pkgver")
sha512sums=('00595551da829a3e5e0a15cc2db703cfe582f47849d2454bb5303b23b310a80cc016ffb3a0ac5d4e96c0d28de77d6dd53a2bfa2d24b86bb1185008eb8b596e98')

build() {
  cd pycryptodome
  python -m build -wn

  cd test_vectors
  python -m build -wn
}

check() {
  cd pycryptodome
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m installer test_vectors/dist/*.whl
  test-env/bin/python -m Crypto.SelfTest
}

package() {
  cd pycryptodome
  python -m installer -d "$pkgdir" dist/*.whl
  install -Dm644 LICENSE.rst -t "$pkgdir"/usr/share/licenses/$pkgname/
  install -Dm644 src/mldsa-native/LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/

  rm -r "$pkgdir"/usr/lib/python3.*/site-packages/Crypto/SelfTest/
}
