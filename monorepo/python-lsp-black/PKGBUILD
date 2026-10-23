# Maintainer: Bruno Pagani <archange@archlinux.org>

pkgname=python-lsp-black
pkgver=2.0.0
pkgrel=6
pkgdesc="python-lsp-server plugin that adds support to black autoformatter, forked from pyls-black"
arch=(any)
url="https://github.com/python-lsp/python-lsp-black"
license=(MIT)
depends=(python-black python-lsp-server)
makedepends=(
    python-build
    python-installer
    python-setuptools
    python-wheel
)
checkdepends=(python-pytest)
source=(${url}/archive/v${pkgver}/${pkgname}-${pkgver}.tar.gz
        https://github.com/python-lsp/python-lsp-black/commit/74d7ae8d.patch
        python-lsp-black-2.0.0-remove-pkg_resources.patch)
sha256sums=('6387584c38e5758431f5d61c269860892f57b7442c0cf9bd712ad64231b9bb1c'
            '8d552419665fb489e3ece681f0313bd386b27e89978a57ad10eebb5cb90349dd'
            '35677a60c630b38c76b5245310fcb5cf8aa36fc80f34dec3a6bdebb6d86e1ea5')

prepare() {
  patch -d $pkgname-$pkgver -p1 < 74d7ae8d.patch # Fix tests with bleck 24.3
  patch -d $pkgname-$pkgver -p1 < python-lsp-black-2.0.0-remove-pkg_resources.patch
}

build() {
  cd ${pkgname}-${pkgver}
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname}-${pkgver}
  python -m venv --system-site-packages test-env
  test-env/bin/python setup.py install --skip-build
  test-env/bin/python -m pytest -vv --color=yes tests \
    -k 'not test_load_config_defaults and not test_load_config_with_skip_options' \
    --deselect tests/test_plugin.py::test_pylsp_format_document_syntax_error \
    --deselect tests/test_plugin.py::test_pylsp_format_range_syntax_error
}

package() {
  cd ${pkgname}-${pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "${pkgdir}"/usr/share/licenses/${pkgname}/
}
