# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=postorius
pkgver=1.3.13
pkgrel=7
pkgdesc="The New Mailman Web UI"
arch=(any)
url="https://gitlab.com/mailman/postorius"
license=(GPL-3.0-or-later)
depends=(
  python
  python-comrak  # via readme_renderer[md]
  python-django
  python-django-allauth  # via python-django-mailman3
  python-django-mailman3
  python-mailmanclient
  python-readme-renderer
)
makedepends=(
  git
  python-build
  python-installer
  python-pdm-backend
)
checkdepends=(
  mailman3
  python-beautifulsoup4
  python-isort
  python-pytest
  python-pytest-django
  python-vcrpy
)
optdepends=(
  'python-django-debug-toolbar: for debugging'
  'python-django-gravatar: for gravatar support'
  'mailman3: for configuring a local mailman instance'
  'uwsgi-plugin-python: for running inside uwsgi'
)
replaces=(python-django-postorius)
backup=(
  etc/uwsgi/${pkgname}.ini
  etc/webapps/$pkgname/settings_local.py
  etc/webapps/$pkgname/urls.py
)
install=$pkgname.install
source=(
  "git+$url?signed#tag=v$pkgver"
  $pkgname-1.3.10-settings.patch
  $pkgname.sysusers
  $pkgname.tmpfiles
  $pkgname.uwsgi
  LICENSE
)
sha512sums=('d0c88d34198840b52aef60ad08a484f8b09392a412618438710fc6e34e93bb3a2059347b5fd0eebdb2319cb669b35dfb2c4d8cc8b229b3f82799f5cfb0d5a376'
            'e465bb865da62e8aaaed5b16361d696ca41775ace5c9b3ca5fbb8e03b1c6f0a1915670f13ee275608e317f4e1a06f1afcc84847c712395e42d436d68fd8c5350'
            '37b43ad58762e098e32226d901be0239b8b66b5fa651bf17462bd05d19000ae2dfd4ca3680cd32cd014ca421bae9ee189f22b803c326c0dad1de774f7b9d51d5'
            '44925c7ce6d1c13bd5bdd2cb449fa978b8eaadaac7ecff74075c23a13478decc4acd1f2cd0399bf55c312d6a08985a3e80f6e7a6e85be9e3dcb93ffb526ca9fb'
            'c042f00d0d8dc72dadd19023152e55246c4d7a4b9d53bca3a05e7b8cb38cac5116417440d4ae82350820c7cb0cd414dcce45fa6af9be19296c4ebd0862dba896'
            'a33658d9271e5c537ccd41bf540b463ad2a5eca4a060c80486ff42a736f0aa042d10436e7177c34d792177cb11285243dee1f31c4df54fb0bfaabbc306406930')
b2sums=('91c8931dc93784222696883d5a5aff49af44904acdafbe65882a5af4d0fb4a138d6f5c413b3a761f665e9e05cbd513b3305010077d786cedd5d28a7a2c7c5d7e'
        '80533bc0e9208c5b97f952e29fd46158d5133a2cb15f64ae8b61195d6c08c67d069b91b04acb04713ca9739d37fbe7808f8d20e45d7a8af65e16cca210220450'
        '9465a19f55b040449dcb636be314c8eb237a55eba3666e1a7c526f3de684da328ffdd717c49c02f1bfd0042c4760f44ddd94e42717cf203793a70a141f1024ef'
        '502d59cb0009b7cdf108a1611d17470d1a73c2d06bf6d9509799b63d4bbb4bb0a2cb71a39eddb0dc5990ee6489625174de6ce2d7e04e77a601c90a6da6d69c84'
        '4b80d60c88d46edd0e6e61ad2ff37ba3568c6c56a2abd449f903f7302aeec4a3bc91bbd3ee846731513fa1906df2e3d7d1c52b75a55214051fb96b048f9b4999'
        'a29664104e1ee73ca0aee1d633e9095d92a57c92787f8d8740bdb7211ba3205782ed8677f539bdb8cae3dd75a3694be3132e185fa3fc4b3f401e1f88eb776101')
validpgpkeys=('541EA0448453394FF77A0ECC9D9B2BA061D0A67C') # Abhilash Raj <raj.abhilash1@gmail.com>

pkgver() {
  cd $pkgname
  git describe | sed 's/\([^-]*-g\)/r\1/;s/v//g;s/-/./g'
}

prepare() {
  # Fix XSS issue: https://www.openwall.com/lists/oss-security/2026/05/07/3
  git -C $pkgname cherry-pick -n 8d00a3c317729f37435bdbd27170f630e341f29e
  # become FHS compliant and disable the use of debug
  patch -d $pkgname -p1 -i ../$pkgname-1.3.10-settings.patch

  # https://gitlab.com/mailman/postorius/-/issues/599
  sed '/example_project/d' -i $pkgname/pyproject.toml

  touch settings_local.py
}

build() {
  cd $pkgname
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
  )

  cd $pkgname
  pytest "${pytest_options[@]}"
}

package() {
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl

  # remove tests
  rm -frv "$pkgdir/$site_packages/$pkgname/tests/"

  # django project
  install -vDm 644 example_project/{__init__,manage,settings,wsgi}.py -t "$pkgdir/usr/share/webapps/$pkgname"
  # symlink locale and static dirs to state dir
  ln -svf /var/lib//$pkgname/locale/ "$pkgdir/usr/share/webapps/$pkgname"
  ln -svf /var/lib/$pkgname/static/ "$pkgdir/usr/share/webapps/$pkgname"
  # state dir
  install -vdm 750 "$pkgdir/var/lib/$pkgname/data"
  install -vdm 755 "$pkgdir/var/lib/$pkgname/"{locale,static}
  # log dir
  install -vdm 750 "$pkgdir/var/log/$pkgname"
  # config
  install -vDm 640 ../settings_local.py example_project/urls.py -t "$pkgdir/etc/webapps/$pkgname"
  ln -svf /etc/webapps/$pkgname/settings_local.py "$pkgdir/usr/share/webapps/$pkgname/settings_local.py"
  ln -svf /etc/webapps/$pkgname/urls.py "$pkgdir/usr/share/webapps/$pkgname/urls.py"
  # uwsgi
  install -vDm 644 ../$pkgname.uwsgi "$pkgdir/etc/uwsgi/$pkgname.ini"
  # tmpfiles.d
  install -vDm 644 ../$pkgname.tmpfiles "$pkgdir/usr/lib/tmpfiles.d/$pkgname.conf"
  # sysusers.d
  install -vDm 644 ../$pkgname.sysusers "$pkgdir/usr/lib/sysusers.d/$pkgname.conf"
  install -vDm 644 ../LICENSE "$pkgdir/usr/share/licenses/$pkgname/0BSD.txt"
}
