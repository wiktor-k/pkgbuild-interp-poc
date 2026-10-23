# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Alex Brinister <alex_brinister at yahoo dot com>

pkgbase=python-cheetah3
pkgname=(python-cheetah3 python-cheetah3-docs)
pkgver=3.4.0.post5
pkgrel=1
pkgdesc='A Python powered template engine and code generator'
arch=(x86_64)
url='https://cheetahtemplate.org'
license=(MIT CC-PDDC)
makedepends=(
  git
  python
  python-legacy-cgi
  python-build
  python-installer
  python-setuptools
  python-wheel
  python-sphinx
)
checkdepends=(python-markdown python-pygments)
source=(
  "$pkgbase::git+https://github.com/CheetahTemplate3/cheetah3#tag=$pkgver"
  cheetah.1
)
sha512sums=('f9402746c3fca7073c480e7030ecfcb655a910335eff529c7717e7e8a42991c8e4b55dab3a4b761361a6faa04a7ac3f5c81b61f39705c408f935a5195893b19b'
            'ae100dee3893c120db8f6544b04db434b412bf0123f7feb40cca15e1866e0e6940a3d74d9baf87b04ebae7314f564929ebc7f34ae5eb72418200f7b01358235e')
b2sums=('e92eadf74e0521bd5f1e4f082ef6d02ced4e66ff4f1d0839c372454358d7b26dd5183cd7bbff749777d8da5cd94df740222aa91f5e72f25086f7607dd249231d'
        '59c0689df18640aea8e671cd267741d1bff4b3954147a34ac71c67ec8399c08460976ecabdce7c7b866bd895c8963b6f03059e496eb7dfa2710d99deba78503d')

prepare() {
  cd "$pkgname"

  # fix: remove deprecated pkg_resources
  # https://github.com/CheetahTemplate3/cheetah3/pull/72
  git cherry-pick --no-commit 7d97ad02420f7c121705b25a24be1926d2a85b6a
}

build() {
  cd "$pkgbase"

  python -m build --wheel --no-isolation

  # generate documentation
  PYTHONPATH="$PWD:$PYTHONPATH" make -C docs html
}

check() {
  cd "$pkgbase"

  # temporary install
  python -m installer --destdir="$(pwd)/tmp" dist/*.whl
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  export PYTHONPATH="$(pwd)/tmp/$site_packages"
  export PATH="$(pwd)/tmp/usr/bin:$PATH"

  python Cheetah/Tests/Test.py
}

package_python-cheetah3() {
  depends=(glibc python python-legacy-cgi)
  optdepends=(
    'python-markdown: for markdown filter'
    'python-memcache: for caching'
    'python-pygments: for code highlight filter'
    'python-cheetah3-docs: documentation'
  )

  cd "$pkgbase"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # man pages
  install -vDm644 -t "$pkgdir/usr/share/man/man1" "$srcdir/cheetah.1"
  ln -sf cheetah.1.gz "$pkgdir/usr/share/man/man1/cheetah-compile.1.gz"
  ln -sf cheetah.1.gz "$pkgdir/usr/share/man/man1/cheetah-analyze.1.gz"

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgbase" LICENSE
}

package_python-cheetah3-docs() {
  pkgdesc+=' (documentation)'

  cd "$pkgbase"

  # documentation
  install -vDm644 -t "$pkgdir/usr/share/doc/$pkgbase" README.rst
  cp -vr docs/_build/html "$pkgdir/usr/share/doc/$pkgbase"

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
