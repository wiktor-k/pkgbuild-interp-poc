# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Guillaume Horel <guillaume.horel@gmail.com>

pkgname=python-psycopg-pool
# source is shared with python-psycopg package, hence the weird naming
_pkgname=python-psycopg
pkgver=3.3.3
pkgrel=1
pkgdesc='Connection pool for psycopg'
arch=(any)
url='https://www.psycopg.org/psycopg3/'
license=(LGPL-3.0-only)
depends=(python python-psycopg)
makedepends=(
  git
  python-build
  python-installer
  python-wheel
  python-setuptools
)
source=("$_pkgname::git+https://github.com/psycopg/psycopg.git#tag=pool-$pkgver")
sha512sums=('11af941cf43e3ca0f1310ee026fbf33a71d3757c0bf97b33af406597f3119f080822640a354af6f7b63dd98efe346e220d1ae98c23846fbbe1277cdf28838493')
b2sums=('64ae3ced21d75049358ee3d579da111ca379c35948252c244c4f0626d71d25f435c27ce29dd01553be1017289a00884c6fe02c04b0ab80eefc53b77972b0b065')

build(){
  cd "$_pkgname/psycopg_pool"

  python -m build --wheel --no-isolation
}

package(){
  cd "$_pkgname/psycopg_pool"

  python -m installer --destdir="$pkgdir" dist/*.whl
}
