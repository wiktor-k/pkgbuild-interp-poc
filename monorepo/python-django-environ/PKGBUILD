# Maintainer: David Runge <dvzrv@archlinux.org>

_name=django-environ
pkgname=python-django-environ
pkgver=0.14.0
pkgrel=1
pkgdesc="Utilize 12factor inspired environment variables to configure your Django application"
arch=(any)
url="https://github.com/joke2k/django-environ"
license=(MIT)
depends=(
  python
  python-django
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pytest)
optdepends=('python-simplejson: for an alternative JSON implementation')
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz fix-postgres_cluster-test-fixture.patch)
sha512sums=('e84c8fe3106f58c195076b37d8c18240363e019e861e5719ddbc0cae80fd420ab301d8470d9dfdb33672f436f0df429258cc67dd8c096ad22f5fea94784567c7'
            '533356ebb382bf8e6d0a950a5e8879b48f2023d9a3a0950bd3195fa97025211846c82d1fe762802f7eab02e1a90d6dd3a8faea06e3ac96617fc2642885d4ab77')
b2sums=('216edc33d00cebaaa36f13a1dcbbcfb128fe0b85d06eb638bd68c6b40778e56a77d4c449c04c3a51cea439760eeb6782fac469751aafaa6748b2068b2dc88c63'
        'a76541343a474509fe466dc55038d6dc2e4ace22c60ae59338c005ffe84ad15f32b4638ec2d886bbd5e5460f441e50e57d4fb605ddce0884147d893cabc8b9fe')

prepare() {
  cd $_name-$pkgver
  patch -Np1 -i ${srcdir}/fix-postgres_cluster-test-fixture.patch
}

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $_name-$pkgver
  pytest -vv
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {AUTHORS,CHANGELOG,README}.rst -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 LICENSE.txt -t "$pkgdir/usr/share/licenses/$pkgname/"
}
