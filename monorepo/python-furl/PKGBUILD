# Maintainer: David Runge <dvzrv@archlinux.org>
# Contributor: Husam Bilal <husam212 AT gmail DOT com>

_name=furl
pkgname=python-furl
pkgver=2.1.4
pkgrel=2
pkgdesc="A small Python library that makes parsing and manipulating URLs easy"
arch=('any')
url="https://github.com/gruns/furl"
license=('Unlicense')
depends=('python' 'python-orderedmultidict' 'python-six')
makedepends=('python-build' 'python-installer' 'python-setuptools' 'python-wheel')
checkdepends=('python-pytest')
source=(
  "https://files.pythonhosted.org/packages/source/${_name::1}/${_name}/${_name}-${pkgver}.tar.gz"
  "fix-tests.patch"
)
sha512sums=('cfd160bb0b87c9b8410690836bf9ec106f981a1f939c90dd1c213800085c7de807ebbacef50a6b7a88894cd12c24f6505ddc4cd111f23fa7615f7ed11a02f027'
            '13d0f817d0fa80f32cdda719d800ff6a8911f6d5b9efdd6a92e9c398b1d3e04a36458236e5c55e87c52c6007b04ced12d9ece3826eded58de428e10ce5ca1bf7')
b2sums=('279a3d1da84625878210e3f47cf9757cfbd07dbe1231927d509ab2f698157462c6372e28a2bb8a4e5586538058f1b3484f9c9e6fadde4dfa23906c6fc3e6d76d'
        '4b984fa3c51f21034ccea22e12ac839028b4e2d8e80469867e1a1a40925978dcd2d91dd8d53a8764091dc2dc5d8500a1f7a37e227527a3a9015798c9a722dc92')

prepare() {
  cd "${_name}-${pkgver}"
  patch -Np1 -i ../fix-tests.patch
}

build() {
  cd "${_name}-${pkgver}"
  python -m build --wheel --skip-dependency-check --no-isolation
}

check() {
  cd "${_name}-${pkgver}"
  export PYTHONPATH="build:${PYTHONPATH}"
  # Ignore tests that were failing prior to Python 3.12 rebuild
  pytest -v -k 'not test_hosts and not test_netloc'
}

package() {
  cd "${_name}-${pkgver}"
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -vDm 644 README.md -t "${pkgdir}/usr/share/doc/${pkgname}"

  # Symlink license file
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  install -d "${pkgdir}/usr/share/licenses/${pkgname}"
  ln -s "${site_packages}/${_name}-${pkgver}.dist-info/LICENSE.md" \
    "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE.md"
}
