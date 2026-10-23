# Maintainer: kpcyrd <kpcyrd[at]archlinux[dot]org>

pkgname=python-pylsqpack
_name=pylsqpack
pkgver=0.3.24
pkgrel=1
pkgdesc='Python bindings for ls-qpack'
url='https://github.com/aiortc/pylsqpack'
arch=('x86_64')
license=('BSD-3-Clause')
depends=('python')
makedepends=('python-build' 'python-installer' 'python-setuptools' 'python-wheel')
source=(https://files.pythonhosted.org/packages/source/${_name::1}/$_name/$_name-$pkgver.tar.gz)
b2sums=('48bfa64f309013d40ad4afd898eee6068179441b1f1bc8e113603e46aa73830bd30a3deb8b80969c219e7d92ae30f6eb99a19b4d8d095de4825fb2e35531d2f7')

build() {
  cd "${_name}-${pkgver}"
  python -m build --wheel --no-isolation
}

check() {
  local _site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd ${_name}-${pkgver}
  python -m installer --destdir=test_dir dist/*.whl

  PYTHONPATH="test_dir/$_site_packages:$PYTHONPATH" python -c 'import pylsqpack'
}

package() {
  cd "${_name}-${pkgver}"
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -D LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}"
}

# vim:set ts=2 sw=2 et:
