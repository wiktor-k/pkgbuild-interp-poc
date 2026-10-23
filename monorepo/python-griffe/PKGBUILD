# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgbase=python-griffe
pkgname=(
  python-griffe
  python-griffelib
  python-griffecli
)
pkgver=2.4.0
pkgrel=1
pkgdesc='Signatures for entire Python programs'
arch=(any)
url='https://mkdocstrings.github.io/griffe'
license=(ISC)
makedepends=(
  git
  python-build
  python-installer
  python-hatchling
)
source=(
  "$pkgbase::git+https://github.com/mkdocstrings/griffe#tag=$pkgver"
  simple-version.patch
)
sha512sums=('aa8222ceb1d7049919c2d53f17f4a788b662c1797aec18eb014054b102ae1656e54928a7c578d7368f4c9ad5229946a9ec096d5c81d1f065e5b750ff9803f4a7'
            'cfc9a0fb1742711dd82ed612241a8553cd4d65d64aadf1f3f6cc716d2318fb2b9860e1a8ce6be23d76a69909ddea85b85b180cd502f2c8e80a8f54907b9ad2f9')
b2sums=('175ef628aee67b664a90a9b8ce539a94d1cc81318d765ef13d348d44b9c733b3bbab752e9d0d4acf39155f776ca229460f936bfc24b095b7beb559a2ca018eba'
        '45e007f8858f4e18ca0767bf99116e9b69946081543bfdd2395dfd52d7ac16d92721a5763a6dcd02e620005b7ba4b61ebf51ee3a79f94f248bc480716eb5572a')

prepare() {
  cd "$pkgbase"

  # * avoid pulling in an entire dependency tree just so some
  # python project can magically set a simple number
  # * replace get_version() with a simpler implementation
  # * remove pdm-backend dependency
  patch -p1 -i "$srcdir/simple-version.patch"
  sed -e "s/@pkgver@/$pkgver/" -i scripts/get_version.py
}

build() {
  cd "$pkgbase"

  pushd packages/griffelib
  python -m build --wheel --no-isolation
  popd

  pushd packages/griffecli
  python -m build --wheel --no-isolation
  popd

  python -m build --wheel --no-isolation
}

package_python-griffe() {
  depends=(
    python
    python-griffelib
    python-griffecli
  )
  cd "$pkgbase"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}

package_python-griffelib() {
  pkgdesc+=' (library package)'
  depends=(
    python
    python-griffecli
    python-typing_extensions
    # griffelib[pypi] dependencies
    python-platformdirs
    python-pip
    python-wheel
  )

  cd "$pkgbase"

  python -m installer --destdir="$pkgdir" packages/griffelib/dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}

package_python-griffecli() {
  pkgdesc+=' (cli package)'
  depends=(
    python
    python-griffelib
    python-colorama
  )

  cd "$pkgbase"

  python -m installer --destdir="$pkgdir" packages/griffecli/dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
