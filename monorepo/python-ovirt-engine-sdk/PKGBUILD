# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Contributor: Konstantin Shalygin <k0ste@k0ste.ru>

pkgname=python-ovirt-engine-sdk
pkgver=4.6.3
pkgrel=3
pkgdesc='Python access to the oVirt Engine API'
arch=('x86_64')
url='http://ovirt.org/develop/release-management/features/infra/python-sdk'
depends=('python-pycurl' 'python-six' 'libxml2')
makedepends=('python-build' 'python-installer' 'python-setuptools' 'python-wheel')
license=('Apache-2.0')
source=("$pkgname-$pkgver.tar.gz::https://github.com/oVirt/python-ovirt-engine-sdk4/archive/refs/tags/$pkgver.tar.gz")
sha512sums=('37876c2a411f51a2c67d52b9be40d208ae71757275888308ca915c8907d4cbf8ba8aadc2ebeb363a9c210bb6ca39df8f21a5458778828c973c2acde7c33716b6')

build() {
  cd python-ovirt-engine-sdk4-$pkgver

  # We just run this command for the templating that it does, we don't do rpms in Arch
  .automation/build-srpm.sh || true

  python -m build --wheel --no-isolation
}

package() {
  cd python-ovirt-engine-sdk4-$pkgver

  python -m installer --destdir="$pkgdir" dist/*.whl

  install -Dm644 LICENSE.txt "${pkgdir}"/usr/share/licenses/${pkgname}/LICENSE
}
