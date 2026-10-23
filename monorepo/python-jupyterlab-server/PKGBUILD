# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Contributor: Guillaume Duboc <guilduboc@gmail.com>

_pyname=jupyterlab_server
pkgname=python-${_pyname/_/-}
pkgver=2.28.1
pkgrel=1
pkgdesc='Launch an application built using JupyterLab'
arch=(any)
url='https://jupyter.org/'
license=(BSD-3-Clause)
depends=(jupyter-server
         python
         python-babel
         python-json5
         python-jsonschema
         python-jupyter-core
         python-packaging
         python-requests
         python-ruamel-yaml
         python-tornado
         python-traitlets)
makedepends=(git
             python-build
             python-hatchling
             python-installer)
checkdepends=(python-jupyter-server-terminals
              python-openapi-core
              python-parse
              python-pytest
              python-pytest-jupyter
              python-pytest-timeout
              python-pytest-tornasync
              python-requests-mock
              python-strict-rfc3339)
conflicts=(jupyterlab_server
           python-jupyterlab_server)
provides=(python-jupyterlab_server)
replaces=(jupyterlab_server
          python-jupyterlab_server)
source=(git+https://github.com/jupyterlab/jupyterlab_server#tag=v$pkgver
        openapi-core-0.23.patch)
sha256sums=('3444e77a59aaf5db52ea0f9405c9b0b92460119579ece498562b9ce5919d8a66'
            '944cfb1848f747f166717c252c127920dc1b9721527f8125260809fb33267792')

prepare() {
  cd $_pyname
  patch -p1 -i ../openapi-core-0.23.patch # Support openapi-core 0.23
}

build() {
  cd $_pyname
  python -m build --wheel --no-isolation
}

check() {
  cd $_pyname
# Skip tests that fail with openapi-core>=0.17
  pytest -v -W ignore::DeprecationWarning -k 'not test_translation_api.py and not test_listings_api.py and not test_settings_api.py and not test_themes_api.py and not test_workspaces_api.py and not test_page_config'
}

package() {
  cd $_pyname
  python -m installer --destdir="$pkgdir" dist/*.whl

  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname
}
