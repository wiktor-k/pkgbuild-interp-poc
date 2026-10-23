# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-pipenv-to-requirements
pkgver=0.9.0
pkgrel=12
pkgdesc="Generate requirements[-dev].txt from Pipfile using pipenv"
url="https://github.com/gsemet/pipenv-to-requirements"
license=('MIT')
arch=('any')
depends=('python-pbr' 'python-pipenv')
checkdepends=('python-pyfakefs' 'python-pytest-mock')
makedepends=('git' 'python-build' 'python-installer' 'python-wheel' 'python-setuptools')
source=("git+https://github.com/gsemet/pipenv-to-requirements.git#tag=$pkgver"
)
sha512sums=('34e501cd777308af49347e3ba3b2b01a6da478acb2ccb6f25ba6091348e7143190e2ac7532919124bd860fa775acf2ab75e7071b8fd56eac13b884a632a8ed26')

prepare() {
  cd pipenv-to-requirements
  git cherry-pick -n fa7cdcd56669c141e3b2a56340b198071680da6a
}

build() {
  cd pipenv-to-requirements
  python -m build --wheel --no-isolation
}

check() {
  cd pipenv-to-requirements
  # https://github.com/gsemet/pipenv-to-requirements/pull/26#issuecomment-1370418878
  pytest --ignore pipenv_to_requirements/test_parsing.py
}

package() {
  cd pipenv-to-requirements
  python -m installer --destdir="$pkgdir" dist/*.whl

  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
