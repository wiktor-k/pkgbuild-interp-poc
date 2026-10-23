# Maintainer: David Runge <dvzrv@archlinux.org>
# Contributor: Daniel M. Capella <polyzen@archlinux.org>

pkgname=khard
pkgver=0.22.0
pkgrel=1
pkgdesc='Console address book manager'
arch=(any)
url=https://github.com/lucc/khard
license=(GPL-3.0-only)
depends=(
  python
  python-configobj
  python-ruamel-yaml
  python-vobject
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools-scm
  python-sphinx
  python-sphinx-argparse
  python-sphinx-autoapi
  python-wheel
)
checkdepends=('python-pytest')
optdepends=(
  'diffutils: Using sdiff_khard_wrapper.sh'
  'vdirsyncer: Synchronization of address books with a DAV server'
)
source=(
  "git+$url.git#tag=v$pkgver"
)
b2sums=('cf4505ef4c57154f4f705b94d6971cf4a48467e819a5876749e9e79020bbb252bc6f2eaa76be028947edbbe7764f2cadf0b63c726f1385aeae487d0711f734b5')

build() {
  cd $pkgname
  SETUPTOOLS_SCM_PRETEND_VERSION=$pkgver python -m build --wheel --skip-dependency-check --no-isolation
  make -C doc man
}

check() {
  cd $pkgname
  pytest -v
}

package() {
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
  # additional wrapper script
  install -vD misc/sdiff/sdiff_${pkgname}_wrapper.sh \
    "$pkgdir"/usr/lib/$pkgname/sdiff_${pkgname}_wrapper.sh
  # twinkle integration
  install -vDm 644 misc/twinkle/scripts/*.py \
    -t "$pkgdir"/usr/share/$pkgname/twinkle/scripts/
  install -vDm 644 misc/twinkle/sounds/*.ogg \
    -t "$pkgdir"/usr/share/$pkgname/twinkle/sounds/
  # zsh
  install -vDm 644 misc/zsh/_*$pkgname \
    -t "$pkgdir"/usr/share/zsh/site-functions/
  # docs
  install -vDm 644 {CHANGES,CONTRIBUTING.rst,README.md} \
    -t "$pkgdir"/usr/share/doc/$pkgname/
  # man
  install -vDm 644 doc/build/man/$pkgname.1 -t "$pkgdir"/usr/share/man/man1
  install -vDm 644 doc/build/man/$pkgname.conf.5 -t "$pkgdir"/usr/share/man/man5
  install -vDm 644 $pkgname/data/{config.spec,template.yaml} -t "$pkgdir/$site_packages"/$pkgname/data/
}
