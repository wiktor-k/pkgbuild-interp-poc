# Maintainer: Maxime Gauduin <alucryd@archlinux.org>

pkgname=python-graphql-core
pkgver=3.3.0
pkgrel=1
pkgdesc='GraphQL base implementation for Python'
url=https://github.com/graphql-python/graphql-core
arch=(any)
license=(MIT)
depends=(
  python
  python-typing_extensions
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-uv-build
  python-wheel
)
checkdepends=(
  python-pytest
  python-pytest-asyncio
  python-pytest-benchmark
)
source=(git+https://github.com/graphql-python/graphql-core.git#tag=v${pkgver})
b2sums=('d145d878d30a506bffead005a79badd01969d01057886fbc3cd03901eee09914e72dd5609b4bd1241e979aa88e1ff9ed9aa3b638de9aa1083ac0f1a9a3e7ce79')

prepare() {
  cd graphql-core
  # HACK: workaround older setuptools requirements
  sed -i 's/setuptools>=59,<76/setuptools/' pyproject.toml
}

build() {
  cd graphql-core
  python -m build --wheel --no-isolation
}

check() {
  cd graphql-core
  PYTHONPATH="$PWD/src" pytest -vv -o addopts=''
}

package() {
  python -m installer --destdir="${pkgdir}" graphql-core/dist/*.whl
  install -Dm 644 graphql-core/LICENSE -t "${pkgdir}"/usr/share/licenses/python-graphql-core/
}

# vim: ts=2 sw=2 et:
