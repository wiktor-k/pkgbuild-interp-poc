# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Contributor: Bao Trinh <qubidt at gmail>
# Contributor: Bastien Traverse <firstname at lastname dot email>

pkgname=python-proxmoxer
_name=${pkgname#python-}
pkgver=2.3.0
pkgrel=1
pkgdesc="Python Wrapper for the Proxmox 2.x API (HTTP and SSH)"
arch=('any')
url="https://github.com/proxmoxer/proxmoxer"
license=('MIT')
depends=('python')
optdepends=('python-requests: for https backend'
			'python-paramiko: for ssh_paramiko backend')
makedepends=('python-setuptools' 'python-build' 'python-installer' 'python-wheel')
source=("https://files.pythonhosted.org/packages/source/${_name::1}/${_name}/${_name}-${pkgver}.tar.gz")
sha256sums=('0b0b2067187af1fc6d4257a46ac68c8fd79cc25d2813637cdae7a9a98fbfd11f')

build() {
  cd "${srcdir}/${_name}-${pkgver}"
  # remove tests because they interfere with other system packages
  sed -i "s/, 'tests', 'tests.base'//" setup.py
  python -m build --wheel --no-isolation
}

package() {
  cd "${srcdir}/${_name}-${pkgver}"
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm 644 LICENSE.txt "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}
