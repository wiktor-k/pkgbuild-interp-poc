# Maintainer: Andreas 'Segaja' Schleifer <segaja at archlinux dot org>

_name='gp-libs'
pkgname="python-${_name}"
pkgver=0.0.19
pkgrel=1
pkgdesc='Incubator for pytest and sphinx helpers for git-pull python projects'
arch=('any')
url='https://gp-libs.git-pull.com/'
license=('MIT')
depends=(
  python
  python-docutils
  python-myst-parser
)
makedepends=(
  python-build
  python-hatchling
  python-installer
  python-sphinx
)
checkdepends=(
  python-pytest
  python-pytest-mock
  python-pytest-rerunfailures
)
source=("https://github.com/git-pull/gp-libs/archive/v${pkgver}/${pkgname}-${pkgver}.tar.gz")
sha512sums=('f1c62fc763b6a679a4e706a332b0798aedc69a57b9615c523f2e6157ffcbb69a2b9ab8d9dbcbcfa4aa8be47fb8dcf5e94f70eebad5112d91678212caf62a89ec')
b2sums=('67118ace790078200d27ae20446afd26cf9a1f066d76a7ea8bd2d17ee6eeabba61bbc7cee980a1eb0e89dc768e3bee97bd892e4c0216aa4c1e416850a5d2f97a')

build() {
  cd  "${_name}-${pkgver}"

  python -m build --wheel --skip-dependency-check --no-isolation
}

check() {
  cd  "${_name}-${pkgver}"

  local _site_packages="$(python -c "import site; print(site.getsitepackages()[0])")"

  python -m installer --destdir=test_dir dist/*.whl

  export PYTHONPATH="test_dir${_site_packages}:${PYTHONPATH}"

  pytest -vv tests
}

package() {
  cd  "${_name}-${pkgver}"

  python -m installer --destdir="${pkgdir}" dist/*.whl

  install --verbose -D --mode=0644 LICENSE --target-directory "${pkgdir}/usr/share/licenses/${pkgname}"
  install --verbose -D --mode=0644 *.md CHANGES --target-directory "${pkgdir}/usr/share/doc/${pkgname}"
}

# vim: tabstop=2 shiftwidth=2 expandtab:
