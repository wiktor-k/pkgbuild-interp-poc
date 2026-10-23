# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Maintainer: Robin Candau <antiz@archlinux.org>

pkgname=ansible
pkgver=14.5.0
pkgrel=1
pkgdesc='Official assortment of Ansible collections'
arch=('any')
url='https://pypi.org/project/ansible/'
license=('GPL-3.0-or-later')
depends=('python' 'ansible-core')
provides=('python-ansible_collections')
optdepends=('python-argcomplete: shell completions'
            'python-pyopenssl: openssl modules'
            'python-dnspython: for dig lookup'
            'python-ovirt-engine-sdk: ovirt support'
            'python-boto3: aws_s3 module'
            'python-proxmoxer: proxmox modules'
            'python-openstacksdk: OpenStack module'
            'python-pynetbox: NetBox module'
            'python-ldap: LDAP modules'
            'python-kubernetes: Kubernetes modules'
            'python-keyring: Keyring module'
            'python-github3py: GitHub module'
            'python-lxc: LXC module'
            'python-fqdn: fqdn_valid module'
            'acme-tiny: openssl_certificate module')
makedepends=('python-build' 'python-installer' 'python-wheel' 'python-setuptools')
source=("https://pypi.python.org/packages/source/a/ansible/ansible-${pkgver}.tar.gz"
        "ansible-13.4.0-remove-pkg_resources.patch")
sha512sums=('3bb9403071b5b4976e2ba1df4a2a56194256006b33b7d3dfda5b2b732fa7f21835d8878b9312f695adfaf56ea9bcbab413ce43db74756cd11536266cb092d48a'
            '8055d481ca3f4705822ce22fc3057129150fd647817ac1fad3937382b3af6a3f0379274aff9e9c4c8039266971302c7daea0311ddd8dea4b2ce2150bce06f82d')

prepare() {
  cd "ansible-${pkgver}"
  patch -Np1 -i ../ansible-13.4.0-remove-pkg_resources.patch
}

build() {
  cd "ansible-${pkgver}"
  python -m build --wheel --no-isolation
}

package() {
  cd "ansible-${pkgver}"
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -Dm 644 COPYING "${pkgdir}/usr/share/doc/${pkgname}/COPYING"
}
