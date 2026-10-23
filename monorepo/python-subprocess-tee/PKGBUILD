# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: David Runge <dvzrv@archlinux.org>

pkgname=python-subprocess-tee
pkgver=0.4.2
pkgrel=2
pkgdesc='A subprocess.run that works like tee'
arch=(any)
url='https://github.com/pycontribs/subprocess-tee'
license=(MIT)
depends=(python)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-setuptools-scm
  python-wheel
)
checkdepends=(
  ansible-core
  python-enrich
  python-pytest
  python-pytest-xdist
)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('efd727261c8fffc6f11f55d5e0a976c5204caac774d32cc009660721ab047883fb74634269e3d9c329d6d59835e01f1d118892aefbd24a2fd7d53ba943034d9a')
b2sums=('5d2acf66d6edaf8934f07883574006af70688e8da97563af29759e4be55916155f999ae527ba51e15154154fedff292a9fad5588992dad02cf8bdc4060f9a037')

build() {
  cd "$pkgname"

  SETUPTOOLS_SCM_PRETEND_VERSION=$pkgver python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  local pytest_options=(
    -vv
    --ignore test/test_func.py
  )

  export PYTHONPATH="src:$PYTHONPATH"
  pytest "${pytest_options[@]}"
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # remove tests: https://github.com/pycontribs/subprocess-tee/issues/61
  local _site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  rm -frv "$pkgdir/$_site_packages/subprocess_tee/test/"

  # documentation
  install -vDm644 -t "$pkgdir/usr/share/doc/$pkgname" README.md

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
